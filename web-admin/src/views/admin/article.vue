<script setup>
import {h} from "vue"
import {NButton, NSwitch} from "naive-ui"
import api from "@/api/api.js";

/*数据加载变量*/
const loading = ref(false)
/*按钮触发函数*/
const sendMail = (rowData) => {
  $message.info(`send mail to ${rowData.id}`)
}


/*分页设置*/
const pagination = reactive({
  page: 1,  // 当前页
  pageSize: 12,  // 每页显示的记录数
  pageCount: null,  // 总页数
  prefix({itemCount}) {  // 自定义前缀文本
    return `总数 ${itemCount}`  // 显示 "Total is {itemCount}."
  }
})

/*自动触发*/
onMounted(() => {
  /*加载数据中*/
  loading.value = true

  $loadingBar.start()
  /*获取数据*/
  api.getArt(params).then(res => {
    /*获取数据*/
    data.value = res.data.result.arts
    /*获取数据条数*/
    pagination.itemCount = res.data.result.total
    /*关闭加载*/
    $loadingBar.finish()
    loading.value = false
  })

})

// 查询参数
const params = reactive({
  page: 1,
})

/*触发分页*/
function page(currentPage) {

  if (!loading.value) {
    $loadingBar.start()
    loading.value = true
    /*显示的当前页*/
    pagination.page = currentPage
    /*请求的后端的参数页*/

    params.page = currentPage

    /*请求文章*/
    api.getArt(params).then(res => {
      data.value = res.data.result.arts
      /*关闭加载*/
      $loadingBar.finish()
      loading.value = false
    })


  }
}


// 处理开关状态改变
function handleOpenChange(row) {
  // 更新行数据中的 open 状态
  row.open = row.open === 0 ? 1 : 0;

  // 调用服务器 API 更新状态
  console.log(`Updating open status of item with ID ${row.id} to ${row.open}`)

}

// 自定义轨道样式
function railStyle({checked}) {
  const style = {};
  if (checked) {
    style.background = "rgba(190,150,229,0.8)";

  }
  return style;
}

/*导航分类设置*/
const columns = ref([
  {
    type: "selection"

  },
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
    title: "发布",
    key: "open",
    align: 'center',      // 列内文本居中对齐
    titleAlign: 'center',  // 表头居中对齐
    render(row) {
      return h(
          NSwitch,
          {
            checkedValue: 1,
            uncheckedValue: 0,
            value: row.open,
            size: "medium",
            railStyle: railStyle,
            onClick: () => handleOpenChange(row)
          }
      )
    }
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
    title: "功能",
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
          {default: () => "编辑"}
      )
    }
  }
])
/*列表数据*/
const data = ref([])
/*选中行*/

// 响应式数据
const checkedRowKeys = ref([])

// 行的唯一标识
const rowKey = (row) => row.id
// 处理行选中
// 处理行选中事件
const handleCheck = (rowKeys) => {
  checkedRowKeys.value = rowKeys
  console.log(checkedRowKeys)

}
</script>

<template>


  <n-card class="h-100% min-w-900">
    <!--  筛选 添加框  -->
    <n-card class="h-80px " style="background-color: rgba(250, 250, 252, 1)">
      <n-space class="items-center">
        <div class="w-300px">
          <n-space class="items-center">
            <span>文章</span>
            <n-input id="name" name="name" placeholder="文章" type="text"/>
          </n-space>

        </div>

        <div class="w-300px">
          <n-space class="items-center">
            <span>分类</span>
            <n-input id="name" name="name" placeholder="分类" type="text"/>
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
        :row-key="rowKey"
        class="mt-20"
        min-height="690"
        remote
        size="large"
        @update:page="page"
        @update:checked-row-keys="handleCheck"
    />

  </n-card>

</template>

<style scoped>

</style>