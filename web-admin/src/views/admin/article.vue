<script setup>
import {h} from "vue"
import {NButton} from "naive-ui"

/*数据加载变量*/
const loading = ref(false)
/*按钮触发函数*/
const sendMail = (rowData) => {
  $message.info(`send mail to ${rowData.id}`)
}


/*分页设置*/
const pagination = reactive({
  page: 2,  // 当前页
  pageSize: 12,  // 每页显示的记录数
  pageCount: 8,  // 总页数
  prefix({itemCount}) {  // 自定义前缀文本
    return `总数 ${itemCount}`  // 显示 "Total is {itemCount}."
  }
})

/*自动触发*/
onMounted(() => {
  pagination.itemCount = 60
})


/*触发分页*/
function page(currentPage) {

  if (!loading.value) {
    loading.value = true
    pagination.page = currentPage

    loading.value = false  // 请求完成后取消加载状态
  }
}

/*导航分类设置*/
const columns = ref([
  {
    title: "ID",
    key: "id",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐
  },
  {
    title: "文章",
    key: "title",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐

  },
  {
    title: "描述",
    key: "desc",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "分类",
    key: "sortId",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "标签",
    key: "tags",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "发布",
    key: "open",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center'  // 表头居中对齐
  },
  {
    title: "创建时间",
    key: "created_at",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "更新时间",
    key: "updated_at",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
  },
  {
    title: "Action",
    key: "actions",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NButton,
          {
            size: "small",
            onClick: () => sendMail(row)
          },
          {default: () => "Send Email"}
      )
    }
  }
])
/*列表数据*/
const data = ref([
  {
    "id": 12,
    "title": "委任为通过现场公布v",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:27",
    "updated_at": "2024-09-11 20:47:27"
  },
  {
    "id": 11,
    "title": "是豆腐干山豆根",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:26",
    "updated_at": "2024-09-11 20:47:26"
  }, {
    "id": 10,
    "title": "委任为通过现场公布v",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:27",
    "updated_at": "2024-09-11 20:47:27"
  },
  {
    "id": 9,
    "title": "是豆腐干山豆根",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:26",
    "updated_at": "2024-09-11 20:47:26"
  },
  {
    "id": 8,
    "title": "委任为通过现场公布v",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:27",
    "updated_at": "2024-09-11 20:47:27"
  },
  {
    "id": 7,
    "title": "是豆腐干山豆根",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:26",
    "updated_at": "2024-09-11 20:47:26"
  },
  {
    "id": 6,
    "title": "的回复你你你",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:26",
    "updated_at": "2024-09-11 20:47:26"
  },
  {
    "id": 5,
    "title": "和顺丰单号",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:25",
    "updated_at": "2024-09-11 20:47:25"
  },
  {
    "id": 4,
    "title": "哈佛大学",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:24",
    "updated_at": "2024-09-11 20:47:24"
  },
  {
    "id": 3,
    "title": "撒旦发射点",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:24",
    "updated_at": "2024-09-11 20:47:24"
  },
  {
    "id": 2,
    "title": "发的供货商的风格和",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:23",
    "updated_at": "2024-09-11 20:47:23"
  },
  {
    "id": 1,
    "title": "啊实打实的",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:22",
    "updated_at": "2024-09-11 20:47:22"
  }
])


</script>

<template>


  <n-card class="h-100% ">
    <!--  筛选 添加框  -->
    <n-card class="h-80px " style="background-color: rgba(250, 250, 252, 1)">
      <n-space class="items-center">
        <div class="w-300px">
          <n-space class="items-center">
            <span>文章</span>
            <n-input placeholder="文章" type="text"/>
          </n-space>

        </div>

        <div class="w-300px">
          <n-space class="items-center">
            <span>分类</span>
            <n-input placeholder="分类" type="text"/>
          </n-space>

        </div>


        <n-button type="tertiary">
          搜索
          <template #icon>
            <n-icon
                :class="'i-ep-search'"
                size="15px"
            />
          </template>
        </n-button>
        <n-button type="primary">
          新增
          <template #icon>
            <n-icon
                :class="'i-ep-plus'"
                size="15px"
            />
          </template>
        </n-button>
      </n-space>

    </n-card>

    <!-- 数据表单   -->
    <n-data-table
        :columns=columns
        :data=data
        :loading=loading
        :pagination=pagination
        class="mt-20"
        min-height="690"
        remote
        size="large"
        @update:page="page"

    />

  </n-card>

</template>

<style scoped>

</style>