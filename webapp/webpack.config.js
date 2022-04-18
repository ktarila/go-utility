const path = require('path');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const { ESBuildMinifyPlugin } = require('esbuild-loader');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const CopyPlugin = require("copy-webpack-plugin");
const HTMLInlineCSSWebpackPlugin = require("html-inline-css-webpack-plugin").default;
const PurgecssPlugin = require('purgecss-webpack-plugin');
const glob = require('glob-all');
const CopyWebpackPlugin = require('copy-webpack-plugin');

const production = process.env.NODE_ENV === 'production';

const config = {
  entry: {
    app: './src/index'
  },
  output: {
    path: path.resolve('dist'),
    filename: 'js/[name].js',
    chunkFilename: 'js/[name].[chunkhash:3].js'
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: 'esbuild-loader'
      },
      {
        test: /\.s?css$/,
        // exclude: /node_modules/,
        use: [
          MiniCssExtractPlugin.loader,
          { loader: "css-loader", options: { url: true } },
          "postcss-loader"
        ],
      },
      // the file-loader will copy the file and fix the appropriate url
        {
            test: /\.(ttf|eot|svg|woff(2)?)(\?v=[0-9]\.[0-9]\.[0-9])?$/,
            loader: "file-loader",
            options: {
            name: "[name].[ext]",
            outputPath: "webfonts/", 
            publicPath: "../webfonts/"
            }
        },
        {
            test: /\.(jpg|png|gif|ico)$/,
            use: [{
                    loader: 'file-loader',
                    options: {
                        name: '[name].[ext]',
                        outputPath: 'images/',
                        publicPath:'images/'
                    }  
                  }]
        }
    ],
  },
  devServer: {
    static: {
      directory: path.join(__dirname, 'build'),
    },
    watchFiles: ['src/*'],
    compress: true,
    port: 8080,
  },
  watchOptions: {
    aggregateTimeout: 200
  },
  plugins: [
    new MiniCssExtractPlugin({
      filename: 'css/[name].css'
    }),
    new PurgecssPlugin({
      paths: glob.sync([
            path.join(__dirname, 'html/**/*.html'),
            path.join(__dirname, 'src/**/*.js'),
        ]),
        defaultExtractor: (content) => {
            return content.match(/[\w-/:]+(?<!:)/g) || [];
        },
        whitelistPatterns: [/^choices/, '/^fa-/']
    }),
    new CopyWebpackPlugin(
      { 
        patterns: [
          { from: '*', to: 'images', context: "src/images" },
        ]
      }
    )
  ],
  mode: production ? 'production' : 'development',
  stats: production ? 'normal' : 'minimal'
};

if (production) {
  config.optimization = {
    minimize: true,
    minimizer: [
      new ESBuildMinifyPlugin({
        css: true
      })
    ]
  };

  config.plugins.push(new HTMLInlineCSSWebpackPlugin());
}

module.exports = config;