<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="IP地址:" prop="network">
          <el-input v-model="formData.network" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="掩码:" prop="netmask">
          <el-input v-model.number="formData.netmask" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网关:" prop="gateway">
          <el-input v-model="formData.gateway" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="机房:" prop="idcId">
          <el-input v-model.number="formData.idcId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="VLAN:" prop="vlanId">
          <el-input v-model.number="formData.vlanId" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="名称:" prop="name">
          <el-input v-model="formData.name" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标签:" prop="label">
          <el-input v-model="formData.label" :clearable="false" placeholder="请输入" />
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
  createIdcIpSubnet,
  updateIdcIpSubnet,
  findIdcIpSubnet
} from '@/api/idcIpSubnet'

defineOptions({
    name: 'IdcIpSubnetForm'
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
            network: '',
            netmask: 0,
            gateway: '',
            idcId: 0,
            vlanId: 0,
            name: '',
            label: '',
            describe: '',
        })
// 验证规则
const rule = reactive({
               network : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               netmask : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               idcId : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               label : [{
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
      const res = await findIdcIpSubnet({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reidcIpSubnet
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
               res = await createIdcIpSubnet(formData.value)
               break
             case 'update':
               res = await updateIdcIpSubnet(formData.value)
               break
             default:
               res = await createIdcIpSubnet(formData.value)
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
