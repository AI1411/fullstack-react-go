export default {
  backend: {
    output: {
      mode: "split",
      target: "src/api/generated/client.ts",
      schemas: "src/api/generated/model",
      client: "react-query",
      mock: true,
      prettier: {
        semi: false,
        singleQuote: true,
        trailingComma: "es5",
      },
    },
    input: {
      target: "../backend/docs/api/swagger.yaml", // OpenAPI仕様ファイル
    },
  },
}
