PROTO_SRC_DIR := pkg
PROTO_DIRS := $(shell find $(PROTO_SRC_DIR) -type d)
PROTO_FILES := $(foreach dir,$(PROTO_DIRS),$(wildcard $(dir)/*.proto))

ifeq ($(PROTO),)
	PROTOS := $(PROTO_FILES)
else
	PROTOS := $(shell find $(PROTO_SRC_DIR) -type d -exec bash -c 'find "{}" -name $(PROTO)' \;)
endif

.PHONY: protoc
protoc:
	$(foreach proto,$(PROTOS),protoc --go_out=. --go-grpc_out=. -I $(PROTO_SRC_DIR) $(proto);)
	echo "Proto files generated"
