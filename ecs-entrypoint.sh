#!/usr/bin/env sh

# Loop over all environment variables from secrets manager
for key in $(echo $SECRETS_MANAGER_ENVIRONMENT_VARIABLES | jq -r 'keys[]'); do
  # If not already defined in the environment
	if [ -z $(printenv $key) ]; then
	  # Export the environment variable
		export $key=$(echo $SECRETS_MANAGER_ENVIRONMENT_VARIABLES | jq -r .$key)
	fi
done

./main