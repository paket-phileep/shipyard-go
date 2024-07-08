# Project Directory Structure

This README provides an overview of the purpose of each folder in the project.

## Folders

### admin

This folder contains administration-related applications and tools.

### animations

Holds all the animations used within the project.

### apps

Contains the different applications that make up the project.

### caddy

Files related to the http server.

### cli

Command Line Interface tools and scripts.

### client

Client adaptations using microsoft kiota of http swagger.

### components

Reusable UI components used throughout the project.

### composables

This folder contains reusable composition functions. These functions encapsulate logic that can be shared across multiple components, helping to keep the codebase modular and maintainable.

#### What are Composables?

Composables are functions that leverage the Composition API in frameworks like Vue.js or React. They allow developers to extract and reuse stateful logic outside of components, making it easier to share and manage code. For example, a composable might handle form validation, data fetching, or state management.

#### Example (Vue 3):

In Vue 3, a composable might look like this:

<pre><div class="dark bg-gray-950 rounded-md border-[0.5px] border-token-border-medium"><div class="flex items-center relative text-token-text-secondary bg-token-main-surface-secondary px-4 py-2 text-xs font-sans justify-between rounded-t-md"><span>javascript</span><div class="flex items-center"><span class="" data-state="closed"><button class="flex gap-1 items-center"><svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" fill="none" viewBox="0 0 24 24" class="icon-sm"><path fill="currentColor" fill-rule="evenodd" d="M7 5a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3h-2v2a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-9a3 3 0 0 1 3-3h2zm2 2h5a3 3 0 0 1 3 3v5h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-9a1 1 0 0 0-1 1zM5 9a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h9a1 1 0 0 0 1-1v-9a1 1 0 0 0-1-1z" clip-rule="evenodd"></path></svg>Copy code</button></span></div></div><div class="overflow-y-auto p-4" dir="ltr"><code class="!whitespace-pre hljs language-javascript">// composables/useFetch.js
import { ref, onMounted } from 'vue';

export function useFetch(url) {
  const data = ref(null);
  const error = ref(null);
  const loading = ref(true);

  const fetchData = async () => {
    loading.value = true;
    try {
      const response = await fetch(url);
      data.value = await response.json();
    } catch (err) {
      error.value = err;
    } finally {
      loading.value = false;
    }
  };

  onMounted(fetchData);

  return { data, error, loading };
}
</code></div></div></pre>

#### Example (React):

In React, a similar hook might look like this:
composables/useFetch.js
import { useState, useEffect } from 'react';

export function useFetch(url) {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const fetchData = async () => {
      setLoading(true);
      try {
        const response = await fetch(url);
        const result = await response.json();
        setData(result);
      } catch (err) {
        setError(err);
      } finally {
        setLoading(false);
      }
    };

    fetchData();
  }, [url]);

  return { data, error, loading };
}
`</code></div>``</div></pre>`

### Benefits of Using Composables:

1. **Reusability** : Encapsulate logic in functions that can be reused across multiple components.
2. **Modularity** : Keep your codebase organized by separating concerns.
3. **Testability** : Easier to test isolated functions than components with complex logic.
4. **Maintainability** : Reduce code duplication and make your codebase easier to maintain.

### Usage:

In your components, you can simply import and use these composables to manage specific pieces of logic. For example, in a Vue component:

import { useFetch } from '../composables/useFetch';

export default {
  setup() {
    const { data, error, loading } = useFetch('https://api.example.com/data');

    return { data, error, loading };
  },
};`</code></div>``</div></pre>`

And in a React component:

import React from 'react';
import { useFetch } from '../composables/useFetch';

const MyComponent = () => {
  const { data, error, loading } = useFetch('https://api.example.com/data');

  if (loading) return `<span class="xml"><div>`Loading...`</div>`;
  if (error) return `<span class="xml"><div>`Error: {error.message}`</div>`;

  return `<span class="xml"><div>`{JSON.stringify(data)}`</div>`;
};

export default MyComponent;
`</code></div>``</div></pre>`

By organizing your reusable logic in the** **`composables` folder, you can streamline the development process and create a more maintainable codebase.

### config

Configuration files for the project.

### constants

Constant values used across the project.

### content

Static content files, such as text and media.

### controller

Controller logic for handling requests and responses.

### core

Core functionalities and services.

### coverage

Code coverage reports.

### database

Database schema, migrations, and seed data.

### debug

Debugging tools and scripts.

### documentation

Project documentation and guides.

### download

Files available for download.

### example

Example files and templates.

### fonts

Custom fonts used in the project.

### functions

Utility functions and helpers.

### hooks

Custom React or Vue hooks.

### images

Image assets used in the project.

### lang

Localization and language files.

### math

Contains FAQ related materials and documents.

### middleware

Middleware functions for handling requests.

### mocks

Mock data and services for testing.

### out

Output directory for build artifacts.

### packages

Contains different packages or modules.

### plugins

Plugins used in the project.

### prisma

Prisma ORM schema and client.

### queue

Queue management and processing files.

### router

Routing logic for the project.

### scripts

Various scripts for development and deployment.

### server

Server-side code and applications.

### source

Source code for the project.

### storybook

Storybook configuration and stories.

### styles

Global and component-specific styles.

### temp

Temporary files and directories.

### templates

Template files used for generating content.

### tests

Unit, integration, and end-to-end tests.

### types

Type definitions and interfaces.

### utils

Utility functions and classes.

### vid

Video assets used in the project.

### views

View templates and components.

### web

Web-related code and assets.
