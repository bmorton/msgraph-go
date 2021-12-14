# msgraph-go

This is an experimental library for demoing OpenAPI code generation for Microsoft Graph in Go.  I haven't vetted this for it's usefulness yet, but the entire thing is a bit unweildy for a single client.

## Generating a client

```sh
make setup
make generate
```

## Caveat

Out of the box, I couldn't get a client generated and it crashed with this error.  If I remove the `/compliance` section from `openapi.yaml`, as I have done in [`openapi.fixed.yaml`](openapi.fixed.yaml)

```
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/bmorton/Code/msgraph-go/msgraph/docs/CommunicationsOnlineMeetingApi.md
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/bmorton/Code/msgraph-go/msgraph/api_communications_presence.go
[main] INFO  o.o.codegen.TemplateManager - writing file /Users/bmorton/Code/msgraph-go/msgraph/docs/CommunicationsPresenceApi.md
Exception in thread "main" java.lang.RuntimeException: Could not generate api file for 'ComplianceCompliance'
	at org.openapitools.codegen.DefaultGenerator.generateApis(DefaultGenerator.java:678)
	at org.openapitools.codegen.DefaultGenerator.generate(DefaultGenerator.java:903)
	at org.openapitools.codegen.cmd.Generate.execute(Generate.java:441)
	at org.openapitools.codegen.cmd.OpenApiGeneratorCommand.run(OpenApiGeneratorCommand.java:32)
	at org.openapitools.codegen.OpenAPIGenerator.main(OpenAPIGenerator.java:66)
Caused by: java.lang.NullPointerException: Cannot read field "isArray" because "codegenProperty" is null
	at org.openapitools.codegen.languages.GoClientCodegen.constructExampleCode(GoClientCodegen.java:541)
	at org.openapitools.codegen.languages.GoClientCodegen.constructExampleCode(GoClientCodegen.java:492)
	at org.openapitools.codegen.languages.GoClientCodegen.postProcessOperationsWithModels(GoClientCodegen.java:460)
	at org.openapitools.codegen.DefaultGenerator.processOperations(DefaultGenerator.java:1220)
	at org.openapitools.codegen.DefaultGenerator.generateApis(DefaultGenerator.java:581)
	... 4 more
error Command failed with exit code 1.
info Visit https://yarnpkg.com/en/docs/cli/run for documentation about this command.
```
