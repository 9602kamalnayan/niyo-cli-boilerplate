# $ServiceName

## Table of Contents

- [Installation](#installation)
- [Steps](#steps)

## Installation

```bash
cd server && npm install
```

## Steps

### 1. Add your manager to the devops, infosec mails for approval.

### 2. Once devops and infosec have created/shared the following :-

- a. UAT OIDC Role (Devops)
- b. PROD OIDC Role (Devops)
- c. Bitbucket repo (Infosec)
- d. Sonar Project (Infosec)
- e. Snyk Token(Infosec)
- f. ECR Repo (Devops)

### 3. Replace in the following files with the details provided above

- a. bitbucket-pipelines.yml (Mandatory)

  - Replace UAT_OIDC with the role provided by devops. (Ex. $NIYO_UAT_BITBUCKET_OIDC_ECR_NINJAFACTORY-2 || arn:aws:iam::914620821736:role/ NIYO_UAT_BITBUCKET_OIDC_ECR_NINJAFACTORY-2)
  - Replace PROD_OIDC with the role provided by devops. (Ex. $$PROD_X2_BITBUCKET_OIDC_ECR_NINJAFACTORY-2 || arn:aws:iam::914620821736:role/$PROD_X2_BITBUCKET_OIDC_ECR_NINJAFACTORY)
  - Replace AWS_AWS_ECR_REPOSITORYSITORY with the ecr repo provided by devops. (Ex. 287726214764.dkr.ecr.ap-south-1.amazonaws.com/niyox2/x2-core-algo-risk-analysis-backend)
  - Replace SNYK_TOKEN with the token name provided by infosec. (Ex. SNYK_Token_X2 || "6dwubd34eek02323" )

- b. .kube/deployment/\*-deployment.yml (Mandatory)

  - Replace AWS_ECR_REPOSITORYSITORY with the ecr repo provided by devops. (Ex. $$PROD_X2_BITBUCKET_OIDC_ECR_NINJAFACTORY-2 || arn:aws:iam::914620821736:role/$PROD_X2_BITBUCKET_OIDC_ECR_NINJAFACTORY)

- c. ./kube/config/\*-config.yml (Mandatory)

  - In the bitbucket-pipelines.yml under the deployments steps you can there is one line



  ```bash
  - cp .kube/config/prod-config.yml ~/.kube/config
  ```

  this is used by the bitbucket runner to deploy the code to kubernetes, now all these files .kube/config/\*-config.yml you can either copy from someone else's config file with the services deployed in the same cluster (Cluster is theu unique identifier for these config files.) or simply ask the devops to provide. (Mandatory for deployment)

### 4. Once all the above steps are done, the code is ready to be deployed to kubernetes via bitbucket CI/CD pipelines.