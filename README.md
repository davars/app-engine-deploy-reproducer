# app-engine-deploy-reproducer
Reproduces an issue I'm having deploying a Go App Engine app using the gcloud SDK

1. Look at env.sh to see how GOPATH should be set up (or `$ source env.sh`).  
2. `cd helloworld`
3. `go test mypkgs...` to verify that importing the vendored dep should work in theory
4. `goapp deploy -application $APP_ENGINE_PROJECT  -version goapp-$(git rev-parse --short HEAD)`
5. `gcloud app deploy --project $APP_ENGINE_PROJECT --version gcloud-$(git rev-parse --short HEAD)`
