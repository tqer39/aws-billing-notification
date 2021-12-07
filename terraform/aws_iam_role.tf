resource "aws_iam_role" "billing_notification" {
  provider              = aws.billing-notification
  name                  = "billing-notification"
  force_detach_policies = true
  max_session_duration  = 43200
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Principal = {
          Service = [
            "cloudwatch.amazonaws.com",
            "events.amazonaws.com",
            "lambda.amazonaws.com",
          ]
        }
        Action    = "sts:AssumeRole"
        Condition = {}
      }
    ]
  })
}