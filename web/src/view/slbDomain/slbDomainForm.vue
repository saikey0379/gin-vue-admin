<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="Name:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="负责人:" prop="manager">
          <el-input v-model="formData.manager" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="描述:" prop="describe">
          <el-input v-model="formData.describe" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="集群:" prop="clusterIds">
          <el-input v-model="formData.clusterIds" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="HTTP:" prop="portHttp">
          <el-input v-model.number="formData.portHttp" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="HTTPS:" prop="portHttps">
          <el-input v-model.number="formData.portHttps" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="HTTP2:" prop="http2">
          <el-switch v-model="formData.http2" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
       </el-form-item>
        <el-form-item label="证书:" prop="certId">
          <el-input v-model.number="formData.certId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="自定义配置:" prop="customize">
          <el-input v-model="formData.customize" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Path:" prop="locations">
          <el-input v-model="formData.locations" :clearable="false" placeholder="请输入" />
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
  createSlbDomain,
  updateSlbDomain,
  findSlbDomain
} from '@/api/slbDomain'

defineOptions({
    name: 'SlbDomainForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            name: '',
            manager: '',
            describe: '',
            clusterIds: '',
            portHttp: 0,
            portHttps: 0,
            http2: false,
            certId: 0,
            customize: '',
            locations: '',
        })
// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               describe : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               clusterIds : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               http2 : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               locations : [{
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
      const res = await findSlbDomain({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reslbDomain
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createSlbDomain(formData.value)
               break
             case 'update':
               res = await updateSlbDomain(formData.value)
               break
             default:
               res = await createSlbDomain(formData.value)
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
