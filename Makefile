release-patch:
	latest_tag=$(git tag -l --sort=-creatordate | head -n 1)
	# Extract the version parts

	major=$(echo "$latest_tag" | cut -d'.' -f1)
	minor=$(echo "$latest_tag" | cut -d'.' -f2)
	patch=$(echo "$latest_tag" | cut -d'.' -f3)

	# Increment the patch version
	new_patch=$((patch + 1))

	# Combine the new version
	new_version="${major}.${minor}.${new_patch}"

	# Release with tags
	git tag -am "[release] release version ${new_version}" ${new_version}
	git push --tags

	echo "The version ${new_version} has been released successfully"

.PHONY: release-patch
