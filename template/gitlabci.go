package template

var (
	GitLabCiSRV = `include:
  - https://git.5th.im/open-source/ci-template/raw/master/templates/deployment-gitlab.yml
  - https://git.5th.im/open-source/ci-template/raw/release/templates/ecr-image-push.yml
  - https://git.5th.im/open-source/ci-template/-/raw/master/templates/docker-login.yaml

variables:
  APP_NAME: module {{.Dir}}

image: docker.longbridge-inc.com/long-bridge-core-system/core/golang-builder

precheck:
  stage: precheck
  except:
    - master
    - pre-master
  script:
    - go fmt ./...
    - go vet ./...

after_script:
  - echo "CommitInfo - $CI_COMMIT_TITLE | $CI_COMMIT_SHA"

stages:
  - precheck
  - build
  - push
  - deploy

base_builder:
  extends: .docker-login
  stage: build
  when: manual
  script:
    - cd $CI_PROJECT_DIR
    - docker build -t docker.longbridge-inc.com/${CI_PROJECT_PATH}/builder -f Dockerfile.builder .
    - docker push docker.longbridge-inc.com/${CI_PROJECT_PATH}/builder

.image: &create_image
  extends: .docker-login
  stage: build
  variables:
    ENDPOINT: "guide"
  script:
    - cd $CI_PROJECT_DIR
    - docker build -t docker.longbridge-inc.com/$CI_PROJECT_PATH:$CI_COMMIT_SHA -f ./cmd/${ENDPOINT}/Dockerfile .
    - docker push docker.longbridge-inc.com/$CI_PROJECT_PATH:$CI_COMMIT_SHA

push_devel:
  <<: *create_image
  stage: build
  only:
    - devel
    - devel2

push_canary:
  <<: *create_image
  stage: build
  only:
    - canary
    - eks-canary

push_bj:
  <<: *create_image
  stage: build
  only:
    - pre-master
    - master

push_canary_hk_image:
  extends: .push_template
  stage: push
  variables:
    IMAGE_HOST: docker.longbridge-inc.com
    ENV: canary-hk
    IMAGE_NAME: $CI_PROJECT_PATH
  only:
    - canary
    - eks-canary

push_prod_bj_image:
  extends:
    - .docker-login
    - .push_template
  stage: push
  only:
    - pre-master
    - master
  variables:
    ENV: prod-bj
    TAG: $CI_COMMIT_SHA
    IMAGE_HOST: docker.longbridge-inc.com
    IMAGE_NAME: $CI_PROJECT_PATH

deploy_devel_hk:
  extends: .deploy_template_v2
  variables:
    ENV: test
    APP_NAME: bi-golang-novice-guide
    TAG: $CI_COMMIT_SHA
  only:
    - devel
    - devel2

deploy_canary_hk:
  extends: .deploy_template_v2
  variables:
    ENV: canary-hk-eks-1
    APP_NAME: bi-golang-novice-guide
    TAG: $CI_COMMIT_SHA
  only:
    - canary
  tags:
    - aws-hk-canary-v1

deploy_canary_hk_eks:
  extends: .deploy_template_v2  #不要修改
  tags:
    - aws-hk-canary-v1          #不要修改
  variables:
    ENV: canary-hk-eks-1        #不要修改
    APP_NAME: bi-golang-novice-guide
    TAG: $CI_COMMIT_SHA
  only:
    - eks-canary

deploy-pre-prod-bj:
  extends: .deploy_template_v2
  variables:
    ENV: prod-bj
    APP_NAME: bi-golang-novice-guide-pre
    TAG: $CI_COMMIT_SHA
  before_script:
    - export NAMESPACE=bi-pre
  only:
    - pre-master
  when: manual

deploy_prod_bj:
  extends: .deploy_template_v2
  variables:
    ENV: prod-bj
    APP_NAME: bi-golang-novice-guide
    TAG: $CI_COMMIT_SHA
  only:
    - master
  when: manual  
  `
)
