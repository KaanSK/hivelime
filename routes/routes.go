package routes

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/kaansk/hivelime/middleware/verifyrequest"
	"github.com/kaansk/hivelime/sublime"
	"github.com/kaansk/hivelime/thehive"
	"github.com/kaansk/hivelime/utils"
)

func SetUpRoutes(app *fiber.App, service utils.Service) {
	sublimeRoutes := app.Group("/sublime")

	if service.Config.SublimeSigningKey != "" {
		sublimeRoutes.Use(
			verifyrequest.New(verifyrequest.Config{
				Secret:     service.Config.SublimeSigningKey,
				Expiration: service.Config.SublimeHMACExpiration}),
		)
	}

	sublimeRoutes.Post("/event", RegisterSublimeSecEvent(service))
}

func RegisterSublimeSecEvent(service utils.Service) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		sublimeEvent := new(sublime.SublimeEvent)
		if err := ctx.BodyParser(sublimeEvent); err != nil {
			return ctx.Status(400).JSON(err.Error())
		}
		tags := []string{}
		mg, _ := service.SublimeClient.GetMessageGroup(sublimeEvent.Data.Message.CanonicalID)
		alert := thehive.NewAlert()
		alert.Type = service.Config.TheHiveAlertType
		alert.Source = "Sublime"
		alert.SourceRef = sublimeEvent.Data.Message.CanonicalID[0:8]
		alert.ExternalLink = fmt.Sprintf("%s/messages/%s", service.Config.SublimeURL, sublimeEvent.Data.Message.CanonicalID)
		alert.Title = fmt.Sprintf("Sublime Detection Alert")

		mailType := []string{}
		if mg.MessageType.Inbound {
			mailType = append(mailType, "inbound")
		}
		if mg.MessageType.Outbound {
			mailType = append(mailType, "outbound")
			for _, recipient := range mg.DataModel.Recipients.To {
				alert.AddObservable("mail", recipient.Email.Email, []string{"Sublime:Recipient"})
			}
			/* 			for _, recipient := range mg.DataModel.Recipients. {
			   				alert.AddObservable("mail", recipient.Email.Email, []string{"Sublime:Recipient"})
			   			}
			   			for _, recipient := range mdm.Recipients.Bcc {
			   				alert.AddObservable("mail", recipient.Email.Email, []string{"Sublime:Recipient"})
			   			} */
		}

		if mg.MessageType.Internal {
			mailType = append(mailType, "internal")
		}

		tags = append(tags, fmt.Sprintf("Sublime:mail-flow=\"%s\"", mailType))

		DescriptionTitle := fmt.Sprintf("# Detection\n\nLink to detection: [Detection Link](%s/messages/%s)\n\n", service.Config.SublimeURL, sublimeEvent.Data.Message.CanonicalID)
		DescriptionBody := fmt.Sprintf("## Flagged Rules:\n")
		if len(mg.FlaggedRules) > 0 {
			for _, flaggedRule := range mg.FlaggedRules {
				DescriptionBody += fmt.Sprintf("* [%s](%s/rules/%s)\n", flaggedRule.RuleMeta.Name, service.Config.SublimeURL, flaggedRule.RuleMeta.ID)
				alert.AddObservable("detection-rule", flaggedRule.RuleMeta.Name, flaggedRule.RuleMeta.Tags)
				tags = append(tags, fmt.Sprintf("Sublime:detection-rule=\"%s\"", flaggedRule.RuleMeta.Name))
			}
		}
		alert.AddObservable("mail", mg.DataModel.Sender.Email.Email, []string{"Sublime:Sender"})
		alert.AddObservable("mail-subject", mg.DataModel.Subject.Subject, []string{})
		for _, ip := range mg.DataModel.Headers.Ips {
			alert.AddObservable("ip", ip.IP, []string{"Sublime:Header"})
		}
		for _, attachment := range mg.DataModel.Attachments {
			nameTag := fmt.Sprintf("Sublime:AttachmentName=%s", attachment.FileName)
			typeTag := fmt.Sprintf("Sublime:AttachmentExtension=%s", attachment.FileExtension)
			alert.AddObservable("hash", attachment.Sha256, []string{nameTag, typeTag})
		}
		for _, link := range mg.DataModel.Body.Links {
			mismatchTag := fmt.Sprintf("Sublime:HrefLinkMismatch=%t", link.HrefURL.Domain.Valid)
			bodyTag := "Sublime:MailBody"
			alert.AddObservable("url", link.HrefURL.URL, []string{mismatchTag, bodyTag})
		}

		alert.Tags = append(service.Config.TheHiveAlertTags, tags...)
		alert.Description = DescriptionTitle + DescriptionBody

		id, err := service.TheHiveClient.CreateAlert(alert)
		if err != nil || id == "" {
			return ctx.Status(501).SendString(fmt.Sprintf("Could not create TheHive Alert: %v", err.Error()))
		}

		return ctx.SendStatus(200)
	}
}
