# read the workflow template
WORKFLOW_TEMPLATE=$(cat .github/deploy-template.yaml)

# iterate each route in services directory
for SERVICE_NAME in $(ls services); do
    echo "generating workflow for services/${SERVICE_NAME}"

    # replace template route placeholder with route name
    WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{SERVICE_NAME}}/${SERVICE_NAME}/g")

    # save workflow to .github/workflows/{SERVICE_NAME}
    echo "${WORKFLOW}" > .github/workflows/${SERVICE_NAME}-service-deploy.yaml
done