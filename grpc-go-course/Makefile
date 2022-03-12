.PHONY: generateCalculator
.PHONY: generateGreet

generateCalculator:
	protoc .\calculator\calculatorpb\calculator.proto --go_out=plugins=grpc:.

generateGreet:
	protoc .\greet\greetpb\greet.proto --go_out=plugins=grpc:.
