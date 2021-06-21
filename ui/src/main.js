import { createApp } from 'vue'
import App from './App.vue'
import { ElButton, ElSelect } from 'element-plus'
import lang from 'element-plus/lib/locale/lang/uz-uz'
import locale from 'element-plus/lib/locale'
import 'dayjs/locale/uz-latn'

locale.use(lang)

const app = createApp(App)
app.use(ElButton)
app.use(ElSelect)
app.mount('#app')
