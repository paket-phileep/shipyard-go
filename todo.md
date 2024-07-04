collate all dockerfiles into one compose with includes √

run proxies for specified repos

install dependenceis using brew √

clone repositories inside of bundles 


marking as private will ignore versioning 


all modules use lerna so that they can be versioned and also marked as private 

login modules more specicially something like a dashboard, carreers module 

all packages also use lerna so that they can be versioned easily and marked as private


lerna is independant of package managers so can be at the root of the monorepo the whole app can also be a yarn workspace and also we can have golang workspaces just configured manually at the root of the battleship aswell 


apps

- applications

components

- libraries even down to ui libraries
- lerna

utilities

- utilities such as loggers ect
- lerna

modules

- plug and play modules for applications and code
- lerna

plugins

- sad
- lerna

lang 

- kiota sdk translations
- lerna
