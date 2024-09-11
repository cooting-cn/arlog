<script setup>
import {h} from "vue"
import {NButton} from "naive-ui"


const sendMail = (rowData) => {
  $message.info(`send mail to ${rowData.name}`);
};

const createColumns = () => [
  {
    title: "文章",
    key: "wz"
  },
  {
    title: "描述",
    key: "mx"
  },
  {
    title: "分类",
    key: "fl"
  },
  {
    title: "标签",
    key: "tags",

  },
  {
    title: "发布",
    key: "fb",

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
    id: 0,
    wz: "John Brown",
    age: 32,
    address: "New York No. 1 Lake Park",
    tags: ["nice", "developer"]
  },
  {
    id: 1,
    wz: "Jim Green",
    age: 42,
    address: "London No. 1 Lake Park",
    tags: ["wow"]
  },
  {
    id: 2,
    wz: "Joe Black",
    age: 32,
    address: "Sidney No. 1 Lake Park",
    tags: ["cool", "teacher"]
  }
];

const columns = createColumns();
const pagination = {
  pageSize: 10
};


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