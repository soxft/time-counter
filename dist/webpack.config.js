const path = require('path');
const webpack = require('webpack');
module.exports = {
    entry: './counter.ts',
    output: {
        filename: 'counter.js',
        path: path.resolve(__dirname, './')
    },
    mode: 'production',
    module: {
        rules: [{
            test: /\.tsx?$/,
            use: 'ts-loader',
            exclude: /node_modules/
        }]
    },
    resolve: {
        extensions: ['.ts']
    },
}