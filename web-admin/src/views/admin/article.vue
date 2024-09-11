<script setup>
import {h} from "vue"
import {NButton} from "naive-ui"


const sendMail = (rowData) => {
  $message.info(`send mail to ${rowData.name}`);
};

const columns = [
  {
    title: "文章",
    key: "title"
  },
  {
    title: "描述",
    key: "desc"
  },
  {
    title: "分类",
    key: "sortId"
  },
  {
    title: "标签",
    key: "tags",
  },
  {
    title: "发布",
    key: "open",
  },
  {
    title: "创建时间",
    key: "created_at",
  },
  {
    title: "更新时间",
    key: "updated_at",
  },
  {
    title: "Action",
    key: "actions",
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
]

const data = [
  {
    "id": 15,
    "title": "维特鸭肉儒雅而银行",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:31",
    "updated_at": "2024-09-11 20:47:31"
  },
  {
    "id": 14,
    "title": "再多发几张的房间内",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:30",
    "updated_at": "2024-09-11 20:47:30"
  },
  {
    "id": 13,
    "title": "而大家的祝福九年",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:29",
    "updated_at": "2024-09-11 20:47:29"
  },
  {
    "id": 12,
    "title": "的做法是竞争的机会呢",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:29",
    "updated_at": "2024-09-11 20:47:29"
  },
  {
    "id": 11,
    "title": "年纪大加拿大这种年纪",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:28",
    "updated_at": "2024-09-11 20:47:28"
  },
  {
    "id": 10,
    "title": "自动返回指定回复",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:28",
    "updated_at": "2024-09-11 20:47:28"
  },
  {
    "id": 9,
    "title": "撒旦发生的故事",
    "desc": "123123123",
    "content": "123333333333333333",
    "img": "",
    "open": 0,
    "sortId": 0,
    "created_at": "2024-09-11 20:47:27",
    "updated_at": "2024-09-11 20:47:27"
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
  }
]


const pagination = {
  pageSize: 20,
  page-sizes: 20
}


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
    <n-data-table :columns="columns"
                  :data="data"
                  :pagination="pagination"
                  class="mt-30"
                  size="large"
    />

  </n-card>

</template>

<style scoped>

</style>