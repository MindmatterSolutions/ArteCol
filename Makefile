# PROJECT VARIABLES
HOST_PROJECT_ROOT := $(CURDIR)
PROJECT_BUILD_DIR := $(HOST_PROJECT_ROOT)/build

# INCLUDES, EXPORTS, AND VARIABLES
include $(PROJECT_BUILD_DIR)/Makefile.core.mk
-include $(PROJECT_BUILD_DIR)/Makefile.dev.mk

.PHONY: stat

stat: ## Return the project info
	@echo $(value PROJECT_BUILD_DIR)


