resource "aws_iam_role_policy_attachment" "billing_notification_AWSLambdaBasicExecutionRole" {
  provider   = aws.billing-notification
  role       = aws_iam_role.billing_notification.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
  depends_on = [
    aws_iam_role.billing_notification,
  ]
}