module github.com/Uptycs/cloudquery

go 1.24

toolchain go1.24.3

require (
	cloud.google.com/go/storage v1.18.2
	github.com/Azure/azure-sdk-for-go v60.1.0+incompatible
	github.com/Azure/azure-storage-blob-go v0.14.0
	github.com/Azure/go-autorest/autorest v0.11.19
	github.com/Azure/go-autorest/autorest/adal v0.9.18 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.5.10
	github.com/Uptycs/basequery-go v0.8.0
	github.com/aws/aws-sdk-go-v2 v1.41.9
	github.com/aws/aws-sdk-go-v2/config v1.32.20
	github.com/aws/aws-sdk-go-v2/credentials v1.19.19
	github.com/aws/aws-sdk-go-v2/service/acm v1.39.2
	github.com/aws/aws-sdk-go-v2/service/apigateway v1.40.2
	github.com/aws/aws-sdk-go-v2/service/bedrockagent v1.0.0
	github.com/aws/aws-sdk-go-v2/service/cloudformation v1.71.13
	github.com/aws/aws-sdk-go-v2/service/cloudtrail v1.56.0
	github.com/aws/aws-sdk-go-v2/service/cloudwatch v1.57.2
	github.com/aws/aws-sdk-go-v2/service/cloudwatchevents v1.33.0
	github.com/aws/aws-sdk-go-v2/service/codecommit v1.34.0
	github.com/aws/aws-sdk-go-v2/service/codedeploy v1.36.0
	github.com/aws/aws-sdk-go-v2/service/codepipeline v1.47.0
	github.com/aws/aws-sdk-go-v2/service/configservice v1.63.0
	github.com/aws/aws-sdk-go-v2/service/directoryservice v1.39.0
	github.com/aws/aws-sdk-go-v2/service/ec2 v1.304.2
	github.com/aws/aws-sdk-go-v2/service/ecr v1.58.0
	github.com/aws/aws-sdk-go-v2/service/ecs v1.82.0
	github.com/aws/aws-sdk-go-v2/service/efs v1.41.18
	github.com/aws/aws-sdk-go-v2/service/eks v1.84.2
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing v1.34.0
	github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2 v1.55.0
	github.com/aws/aws-sdk-go-v2/service/glacier v1.33.0
	github.com/aws/aws-sdk-go-v2/service/guardduty v1.78.2
	github.com/aws/aws-sdk-go-v2/service/iam v1.54.0
	github.com/aws/aws-sdk-go-v2/service/kms v1.53.0
	github.com/aws/aws-sdk-go-v2/service/organizations v1.51.6
	github.com/aws/aws-sdk-go-v2/service/rds v1.118.4
	github.com/aws/aws-sdk-go-v2/service/s3 v1.102.2
	github.com/aws/aws-sdk-go-v2/service/sns v1.39.19
	github.com/aws/aws-sdk-go-v2/service/sqs v1.42.29
	github.com/aws/aws-sdk-go-v2/service/sts v1.42.3
	github.com/aws/aws-sdk-go-v2/service/workspaces v1.68.3
	github.com/fatih/structs v1.1.0
	github.com/patrickmn/go-cache v2.1.0+incompatible
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.8.1
	github.com/stretchr/testify v1.7.0
	google.golang.org/api v0.58.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

require (
	cloud.google.com/go v0.97.0 // indirect
	github.com/Azure/azure-pipeline-go v0.2.3 // indirect
	github.com/Azure/go-autorest v14.2.0+incompatible // indirect
	github.com/Azure/go-autorest/autorest/azure/cli v0.4.2 // indirect
	github.com/Azure/go-autorest/autorest/date v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.4.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.3.1 // indirect
	github.com/Azure/go-autorest/logger v0.2.1 // indirect
	github.com/Azure/go-autorest/tracing v0.6.0 // indirect
	github.com/Microsoft/go-winio v0.4.16 // indirect
	github.com/apache/thrift v0.14.2 // indirect
	github.com/aws/aws-sdk-go-v2/feature/ec2/imds v1.18.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/configsources v1.4.25 // indirect
	github.com/aws/aws-sdk-go-v2/internal/endpoints/v2 v2.7.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/accept-encoding v1.13.10 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/presigned-url v1.13.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/s3shared v1.19.25 // indirect
	github.com/aws/aws-sdk-go-v2/service/sso v1.30.19 // indirect
	github.com/aws/smithy-go v1.26.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dimchansky/utfbom v1.1.1 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/google/uuid v1.2.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/mattn/go-ieproxy v0.0.1 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	go.opencensus.io v0.23.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210503060351-7fd8e65b6420 // indirect
	golang.org/x/oauth2 v0.0.0-20211005180243-6b3c2da341f1 // indirect
	golang.org/x/sys v0.0.0-20210917161153-d61c044b1678 // indirect
	golang.org/x/text v0.3.6 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211016002631-37fc39342514 // indirect
	google.golang.org/grpc v1.40.0 // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
)

require (
	github.com/aws/aws-sdk-go-v2/aws/protocol/eventstream v1.7.11 // indirect
	github.com/aws/aws-sdk-go-v2/internal/v4a v1.4.26 // indirect
	github.com/aws/aws-sdk-go-v2/service/internal/checksum v1.9.18 // indirect
	github.com/aws/aws-sdk-go-v2/service/signin v1.1.1 // indirect
	github.com/aws/aws-sdk-go-v2/service/ssooidc v1.36.2 // indirect
	github.com/golang-jwt/jwt/v4 v4.0.0 // indirect
)
