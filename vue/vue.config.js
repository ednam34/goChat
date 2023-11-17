const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    allowedHosts: [
      'localhost',
      'rayanekaabeche.fr:8080',
      'rayanekaabeche.fr',
    ],
    disableHostCheck: true,
  },
  transpileDependencies: true
})
