const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    allowedHosts: [
      'localhost',
      'http://rayanekaabeche.fr:8080/',
      'http://rayanekaabeche.fr',
    ],
  },
  transpileDependencies: true
})
