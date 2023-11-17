const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  devServer: {
    allowedHosts: [
      'rayanekaabeche.fr', // Ajoutez ici les hôtes autorisés
      'localhost'
    ],
  },
  transpileDependencies: true
})
