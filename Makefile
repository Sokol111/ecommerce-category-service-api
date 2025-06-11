-include makefiles/generate-api.mk

update-makefiles:
	@echo "Updating includes in Makefile..."
	mkdir -p makefiles
	curl -sSL https://raw.githubusercontent.com/Sokol111/ecommerce-makefiles/master/generate-api.mk -o makefiles/generate-api.mk
	@echo "Done."
	