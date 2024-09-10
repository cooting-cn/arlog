import {createDiscreteApi, darkTheme, lightTheme} from "naive-ui"


const themeRef = ref("light");
const configProviderPropsRef = computed(() => ({
    theme: themeRef.value === "light" ? lightTheme : darkTheme
}));

const {message, notification, dialog, loadingBar, modal} = createDiscreteApi(
    ["message", "dialog", "notification", "loadingBar", "modal"],
    {
        configProviderProps: configProviderPropsRef
    }
)

/*注册全局引用*/
window.$loadingBar = loadingBar
window.$notification = notification
window.$message = setupMessage(message)
window.$dialog = dialog


export function setupMessage(NMessage) {
    class Message {
        static instance

        constructor() {
            // 单例模式
            if (Message.instance)
                return Message.instance
            Message.instance = this
            this.message = {}
            this.removeTimer = {}
        }

        removeMessage(key, duration = 5000) {
            this.removeTimer[key] && clearTimeout(this.removeTimer[key])
            this.removeTimer[key] = setTimeout(() => {
                this.message[key]?.destroy()
            }, duration)
        }


        destroy(key, duration = 200) {
            setTimeout(() => {
                this.message[key]?.destroy()
            }, duration)
        }

        showMessage(type, content, option = {}) {
            if (Array.isArray(content)) {
                return content.forEach(msg => NMessage[type](msg, option))
            }

            if (!option.key) {
                return NMessage[type](content, option)
            }

            const currentMessage = this.message[option.key]
            if (currentMessage) {
                currentMessage.type = type
                currentMessage.content = content
            } else {
                this.message[option.key] = NMessage[type](content, {
                    ...option,
                    duration: 0,
                    onAfterLeave: () => {
                        delete this.message[option.key]
                    },
                })
            }
            this.removeMessage(option.key, option.duration)
        }

        loading(content, option) {
            this.showMessage('loading', content, option)
        }

        success(content, option) {
            this.showMessage('success', content, option)
        }

        error(content, option) {
            this.showMessage('error', content, option)
        }

        info(content, option) {
            this.showMessage('info', content, option)
        }

        warning(content, option) {
            this.showMessage('warning', content, option)
        }
    }

    return new Message()
}