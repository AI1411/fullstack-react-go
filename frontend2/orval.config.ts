module.exports = {
  backend: {
    output: {
      mode: "split",
      target: "src/api/client.ts",
      schemas: "src/api/model",
      client: "react-query",
      mock: true,
      prettier: {
        semi: false,
        singleQuote: true,
        trailingComma: 'es5',
      },
    },
    input: {
      target: "../backend/docs/swagger.yaml", // OpenAPI仕様ファイル
    },
  },
}
