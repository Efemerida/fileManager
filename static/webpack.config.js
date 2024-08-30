const MiniCssExtractPlugin = require("mini-css-extract-plugin");
const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = {
  mode: 'development', 
  entry: './src/scripts/main.ts', // точка входа
  output: {
    filename: 'bundle.[contenthash].js',  // имя выходного файла
    path: path.resolve(__dirname, 'resource'), // путь для выходного файла
    clean: true, // очищает папку 
  },
  resolve: {
    extensions: ['.ts', '.js', 'css'], // обрабатываемые расширения
  },
  module: {
    rules: [
      {
        test: /\.ts?$/, // обработка файлов .ts
        use: 'ts-loader',
        exclude: /node_modules/,
      },
      {
        test: /\.css$/, // обработка файлов .css
        use: [MiniCssExtractPlugin.loader, 'css-loader']
      },
    ]
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './src/index.html', // шаблон HTML
    }),
    new MiniCssExtractPlugin({
        filename: 'style.css', // имя выходного файла для CSS
      }),
  ],
};