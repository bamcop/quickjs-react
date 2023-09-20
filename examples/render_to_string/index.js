import React from 'react'
import { renderToString } from 'react-dom/server'
import App from './app'

globalThis.React = React

console.log(renderToString(<App />))
