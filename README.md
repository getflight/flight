# Flight
Flight is a serverless deployment platform for GO apps powered by AWS Lambda.

```shell
$ flight deploy -e staging
09:41:00 deploying to staging
09:41:00 reading configuration
09:41:00 parsing manifest
09:41:00 packaging artifact
09:41:00 saving artifact
09:41:01 polling artifact for upload
09:41:02 uploading artifact
09:41:03 initiating deployment
09:41:05 ensuring artifact exists
09:41:05 updating function
09:41:12 removing previous functions
09:41:12 updating function alias
09:41:12 deployment #1 completed successfully in 12s
```


## Current state
Flight is currently in the alpha phase and while the features provided have undergone testing prior to release, it is possible that users may encounter bugs or disruptions while utilizing the platform. Should any issues arise, we kindly request that you open a support ticket for prompt resolution.


## Getting started
To utilize Flight, it is necessary to first register for an active account, subsequently create an organisation, and finally link your AWS account.

### Account creation
Creating an account on the [Flight dashboard](https://dashboard.getflight.io/register) is a straightforward process. To do so, you will be required to provide your desired username, full name, and password. An activation email will then be sent to the provided email address, allowing you to fully activate your newly created account.

### Creating an organisation
Upon successful activation of your account, you will have access to the [Flight dashboard](https://dashboard.getflight.io/register), where you will be able to create your initial organisation. This organisation serves as a grouping mechanism for all deployed applications and allows for the invitation of additional users for collaboration purposes. 

To initiate the creation of a new organisation, navigate to the "Create organisation" link located on the dashboard's homepage. You will be prompted to provide a unique organisation name and select a cloud provider. At this time, Flight's primary cloud provider is AWS. 

In order to manage deployments, an active AWS IAM user needs to be linked to the organisation. To accomplish this, you will be required to provide an AWS access key, AWS secret key, and the corresponding AWS region. For guidance on creating an IAM user for your AWS account, please refer to the [following documentation](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_create.html).

### CLI Installation
The Flight command line interface (CLI) can be obtained by downloading the binary from the [latest release page](https://github.com/getflight/flight/releases/latest). Binaries for both Windows and Linux systems are available.

### Logging in
Upon successful activation of your account, and creation of an organisation, you will have the ability to log in by utilizing the following command.

```shell
flight login -e your@email.com
```

If multiple organisations are associated with your account, you will be prompted to select one for use. Subsequently, all future deployments will be directed to the selected organisation. 

### Project configuration
The configuration process requires the presence of a flight.yml file at the root of your project. A simple example is provided below.

```yml
# ./flight.yml

name: my-app
trigger: gateway
environments:
  - name: staging
    databases:
      - name: my-db
        driver: mysql
    variables:
      - key: my-key
        value: my-value
```
Initiating a deployment will automatically provision all necessary resources to enable the execution of your application. For further information, please refer to the configuration section.

### Deployment
To deploy your application, execute the deploy command at the root of your project and specify the desired environment for deployment.
```shell
flight deploy -e staging
```

Prior to deployment, it is essential that your application has been built and the executable is present in the root directory of your project. The root directory will be searched for the executable file name specified in the `flight.yml` configuration file. For instance, with the current configuration, the built application with the file name `my-app.exe` for Windows or `my-app` for Linux will be used for deployment.

Once the deployment is completed, all necessary resources will be set up in your AWS account. You will then have the ability to invoke your application in accordance with the configured trigger.

And with that, congratulations on successfully deploying your first application.



## Configuration

This section provides a comprehensive overview of the various configuration options available through the use of the `flight.yml` file.

### Name
```yml
name: my-app
```

Name of your application. The name of your application must be unique within the environment to prevent any potential overwriting during deployment. Additionally, the configured value will be used to locate the executable during deployment.

### Trigger

The desired method of invoking your application, which can be either `gateway` or `queue`. Detailed information on each option can be found in the sections below.

#### Gateway
```yml
trigger: gateway
```

The `gateway` trigger option enables your function to be invoked through HTTP, making it ideal for web-based applications. Utilizing this trigger type during deployment will result in the provisioning of an AWS API Gateway, providing a unique URL for invocation. For an implementation example, please refer to the [following project](https://github.com/getflight/examples/tree/master/basic).

#### Queue

```yml
trigger: queue
```

The `queue` trigger option allows for invocation of your function through queue messaging, making it ideal for long-running, asynchronous backend processes. When deploying using this trigger type, Flight will set up an AWS SQS queue for your function, enabling invocation through messaging. For an implementation example, please refer to the [following project](https://github.com/getflight/examples/tree/master/queue).

### Files
The `files` configuration enables you to specify multiple folders or files to be included in the deployed artifact, making them accessible during execution. For further information, please refer to the [accompanying project](https://github.com/getflight/examples/tree/master/basic).
```yml
files:
  - resources
```

### Environments
The `environments` configuration allows you to specify the environments that you wish to configure. These environments are only established during execution of the deploy command. Each environment is assigned a private network to guarantee the security of your functions, preventing external access. Applications within the same environment are deployed within the same network, allowing for communication between them.

```yml
environments:
  - name: staging
  - name: prod
```

### Databases
The `databases` configuration enables the creation and linking of databases to your application. These databases are established at the environment level, ensuring that they are only accessible to apps deployed within the corresponding environment. Databases are created during deployment. Once created and available, each subsequent deployment will ensure that the databases are properly configured.

```yml
environments:
  - name: staging
    databases:
      - name: my-db
        driver: mysql
```

Currently, two database drivers are supported: `mysql` and `postgresql`. At this time, the versions of the drivers are not configurable, they are fixed and are as follows:

| Driver      | Version                                              |
|-------------|------------------------------------------------------|
| mysql       | Aurora MySQL (compatible with MySQL 5.7.2.08.3)      |
| postgresql  | Aurora PostgreSQL (Compatible with PostgreSQL 11.18) |

Upon deployment of the app, the necessary environment variables will be created and injected into your application, enabling connection to the database. For instance, for the `my-db` database: 

- `DATABASES_MY_DB_DRIVER`
- `DATABASES_MY_DB_HOST`
- `DATABASES_MY_DB_NAME`
- `DATABASES_MY_DB_PASSWORD`
- `DATABASES_MY_DB_PORT`
- `DATABASES_MY_DB_USERNAME`

#### Recovering authentication details
The username and password for your database will be generated during database creation and securely stored in AWS Secrets Manager. To access these authentication details, you'll need to read the environment variables and retrieve the secret using the Amazon Secrets Manager SDK. 

For example, an app deployed in the `staging` environment with a database named `my-db` will have an environment variable injected with the following value: 

`DATABASE_PASSWORD={secret}flight_managed_staging_my_db_password`

- The `{secret}` prefix indicates that the value is stored in AWS Secrets Manage
- The string `flight_managed_staging_my_db_password` represents the key to retrieve the corresponding secret value. 


By using the AWS Secrets Manager SDK, you'll be able to access the password for your database. For further information, refer to the database example in the [following repository](https://github.com/getflight/examples/tree/master/database).

### Variables

The `variables` configuration allows for the injection of environment variables into your app. These variables are established at the environment level, ensuring that they are only injected into apps deployed within the corresponding environment.

```yml
environments:
  - name: staging
    variables:
      - key: my-key-1
        value: my-value-1
      - key: my-key-2
        value: my-value-2
```

The key of each variable is transformed into uppercase snake case format. For instance, `my-key-1` would be transformed to `MY_KEY_1`.

## Examples
Integration examples are available in the [following repository](https://github.com/getflight/examples).