import { defineConfig } from "orval";

export default defineConfig({
  backend: {
    input: {
      target: "../documents/api/openapi.yaml",
    },
    output: {
      mode: "tags-split",
      clean: true,
      target: "./src/gen/queries",
      schemas: "./src/gen/models",
      client: "react-query",
      prettier: true,
      mock: true,
      tsconfig: "tsconfig.json",
      override: {
        query: {
          useQuery: true,
        },
      },
    },
    hooks: {
      afterAllFilesWrite: "npx prettier --write",
    },
  },
});