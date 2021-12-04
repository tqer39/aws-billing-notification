resource "aws_cloudwatch_log_group" "billing_notification" {
  provider = aws.billing-notification
  name     = "/aws/lambda/billing_notification"
}