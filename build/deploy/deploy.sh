#!/bin/bash

gcloud builds submit --config=build/deploy/gateway-service.yaml && \
gcloud builds submit --config=build/deploy/customer-service.yaml && \
gcloud builds submit --config=build/deploy/catalog-service.yaml && \
gcloud builds submit --config=build/deploy/item-service.yaml
