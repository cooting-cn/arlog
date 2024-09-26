import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import {NaiveUiResolver} from 'unplugin-vue-components/resolvers'

// 引入Unocss
import Unocss from 'unocss/vite';
import {presetAttributify, presetIcons, presetUno} from 'unocss'

// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
    // 根据不同的模式设置不同的基础 URL
    const baseUrl = mode === 'pro' ? 'https://arlog.cn' : 'http://127.0.0.1';

    return {
        plugins: [
            vue(),
            Unocss({
                presets: [
                    presetUno(),
                    presetAttributify(),
                    presetIcons(),
                ],
            }),
            AutoImport({
                imports: [
                    'vue', 'vue-router',
                    {
                        'naive-ui': [
                            'useDialog',
                            'useMessage',
                            'useNotification',
                            'useLoadingBar',
                        ],
                    },
                ],
                dts: false,
            }),
            Components({
                resolvers: [NaiveUiResolver()],
            }),
        ],
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url)),
            },
        },
        // Define global variables for the app
        define: {
            baseUrl: JSON.stringify(baseUrl),
        },
    };
})


