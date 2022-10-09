PROTO_DIR := ./proto

protos:
	protoc --proto_path=$(PROTO_DIR) --go_out=pb --go_opt=paths=source_relative \
    proto/*.proto

# during the clean: if there if a pb directory we delete all the generated pb code else we create the pb directory
clean:
	if [ -d pb ]; then rm -f pb/*.pb; else mkdir pb; fi

.PHONY: clean protos