# Assumption

This section explains the assumption on the system
that was not clearly explained on the problem.

## Payment

I'm not sure how the payment works in the system.
Here I assume that user can click "Pay Now" on billing page
on the app and brought to a payment system with billing id.

So to update payment status, this service will
to listen to payment event in a queue service such
as kafka / SQS and evaluate the billing id in it.