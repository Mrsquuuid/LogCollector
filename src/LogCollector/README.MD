Log collector

Logcollector is a lightweight and dynamically managed log collection client. All the log collection tasks are configured and managed through the log collection management side.

By configuring log collection items on the operation and maintenance platform, logagent obtains the log information to be collected from etcd, reads the log information from the business server, and sends it to Kafka. Logtransfer is responsible for reading the log from Kafka, writing it to elasticsearch, and retrieving the log through kibana. The collection of system performance data includes node_ The alarm information is pushed to the alert manager, and finally displayed visually through grafana.