<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="SN:" prop="sn">
          <el-input v-model="formData.sn" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Hostname:" prop="hostname">
          <el-input v-model="formData.hostname" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="IP:" prop="ip">
          <el-input v-model="formData.ip" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="操作系统版本:" prop="versionOs">
          <el-input v-model="formData.versionOs" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内核版本:" prop="versionKernel">
          <el-input v-model="formData.versionKernel" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="厂商:" prop="manufacturer">
          <el-input v-model="formData.manufacturer" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="型号:" prop="modelName">
          <el-input v-model="formData.modelName" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CPU:" prop="cpuSum">
          <el-input v-model.number="formData.cpuSum" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="CPU信息:" prop="cpu">
          <el-input v-model="formData.cpu" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内存:" prop="memorySum">
          <el-input v-model.number="formData.memorySum" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="内存信息:" prop="memory">
          <el-input v-model="formData.memory" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="网卡:" prop="nicInfo">
          <el-input v-model="formData.nicInfo" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="NIC:" prop="nicDevice">
          <el-input v-model="formData.nicDevice" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Gpu:" prop="gpu">
          <el-input v-model="formData.gpu" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="主板:" prop="motherboard">
          <el-input v-model="formData.motherboard" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Raid卡:" prop="raid">
          <el-input v-model="formData.raid" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="硬盘信息:" prop="disk">
          <el-input v-model="formData.disk" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="Server类型:" prop="serverType">
          <el-select v-model="formData.serverType" placeholder="请选择" :clearable="false">
            <el-option v-for="(item,key) in ServerTypeOptions" :key="key" :label="item.label" :value="item.value" />
          </el-select>
       </el-form-item>
        <el-form-item label="Agent版本:" prop="versionAgent">
          <el-input v-model="formData.versionAgent" :clearable="false" placeholder="请输入" />
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
  createServerInfo,
  updateServerInfo,
  findServerInfo
} from '@/api/serverInfo'

defineOptions({
    name: 'ServerInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const ServerTypeOptions = ref([])
const formData = ref({
            sn: '',
            hostname: '',
            ip: '',
            versionOs: '',
            versionKernel: '',
            manufacturer: '',
            modelName: '',
            cpuSum: 0,
            cpu: '',
            memorySum: 0,
            memory: '',
            nicInfo: '',
            nicDevice: '',
            gpu: '',
            motherboard: '',
            raid: '',
            disk: '',
            serverType: undefined,
            versionAgent: '',
        })
// 验证规则
const rule = reactive({
               sn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               hostname : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               ip : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               versionOs : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               versionKernel : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               manufacturer : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               modelName : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cpuSum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               cpu : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               memorySum : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               memory : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nicInfo : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               nicDevice : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               gpu : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               motherboard : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               raid : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               disk : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               serverType : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               versionAgent : [{
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
      const res = await findServerInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reserverInfo
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    ServerTypeOptions.value = await getDictFunc('ServerType')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createServerInfo(formData.value)
               break
             case 'update':
               res = await updateServerInfo(formData.value)
               break
             default:
               res = await createServerInfo(formData.value)
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
