<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="集群ID:" prop="clusterId">
          <el-input v-model.number="formData.clusterId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="域名:" prop="domainId">
          <el-input v-model="formData.domainId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Path:" prop="routeId">
          <el-input v-model="formData.routeId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="类型:" prop="accessType">
          <el-select v-model="formData.accessType" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in AclTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="节点:" prop="nodes">
          <el-input v-model="formData.nodes" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="describe">
          <el-input v-model="formData.describe" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createSlbAccesslist,
  updateSlbAccesslist,
  findSlbAccesslist
} from '@/api/slbAccesslist'

defineOptions({
    name: 'SlbAccesslistForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const AclTypeOptions = ref([])
const formData = ref({
            clusterId: 0,
            domainId: '',
            routeId: '',
            accessType: undefined,
            nodes: '',
            describe: '',
        })
// 验证规则
const rule = reactive({
               clusterId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               domainId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               routeId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               accessType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nodes : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSlbAccesslist({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reslbAccesslist
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    AclTypeOptions.value = await getDictFunc('AclType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSlbAccesslist(formData.value)
               break
             case 'update':
               res = await updateSlbAccesslist(formData.value)
               break
             default:
               res = await createSlbAccesslist(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
