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
        <el-form-item label="集群:" prop="clusters">
          <el-input v-model="formData.clusters" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="记录:" prop="record">
          <el-input v-model="formData.record" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="警报:" prop="alert">
          <el-input v-model="formData.alert" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="表达式:" prop="expr">
          <el-input v-model="formData.expr" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="持续时间:" prop="for">
          <el-input v-model="formData.for" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="保持时间:" prop="keepFiringFor">
          <el-input v-model="formData.keepFiringFor" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="标签:" prop="labels">
          <el-input v-model="formData.labels" :clearable="false" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="注解:" prop="annotations">
          <el-input v-model="formData.annotations" :clearable="false" placeholder="请输入" />
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
  createRuleRecord,
  updateRuleRecord,
  findRuleRecord
} from '@/api/ruleRecord'

defineOptions({
    name: 'RuleRecordForm'
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
            clusters: '',
            record: '',
            alert: '',
            expr: '',
            for: '',
            keepFiringFor: '',
            labels: '',
            annotations: '',
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
               alert : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               expr : [{
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
      const res = await findRuleRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reruleRecord
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
               res = await createRuleRecord(formData.value)
               break
             case 'update':
               res = await updateRuleRecord(formData.value)
               break
             default:
               res = await createRuleRecord(formData.value)
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
