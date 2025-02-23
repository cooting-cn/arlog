import {defineStore} from "pinia";

const blogStore = defineStore('arlog', {
        state: () => ({
            token: "",
            user: "666",
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