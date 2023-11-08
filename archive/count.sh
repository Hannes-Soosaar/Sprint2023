numberOfFolders=$(find . ! -path '*/.*' | wc -l )
numberOfFoldersTimesFive=$((numberOfFolders*5))
printf "\t\vTotal files * 5: %s\v\n" "$numberOfFoldersTimesFive"
