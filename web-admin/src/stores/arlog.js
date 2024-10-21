import {defineStore} from "pinia";

const blogStore = defineStore('arlog', {
        state: () => ({
            token: "",
            user: "",
            user_img: "",
            key: ""
        }),
        persist: {
            enabled: true,
            strategies: [
                {
                    key: 'arlog',
                    storage: localStorage,
                },
            ],
        }
    }
)

export default blogStore