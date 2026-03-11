<template>
  <div>
    <div class="toolbar">
      <el-input v-model="search" placeholder="搜索设备名/编号" clearable style="width:240px">
        <template #prefix><el-icon><Search /></el-icon></template>
      </el-input>
      <el-select v-model="filterStatus" placeholder="状态筛选" clearable style="width:120px">
        <el-option label="在线" value="online" />
        <el-option label="维护" value="maintain" />
        <el-option label="离线" value="offline" />
      </el-select>
      <el-button type="primary" :icon="Plus">新增设备</el-button>
    </div>

    <el-table :data="filteredMachines" stripe border style="width:100%;margin-top:16px" size="small">
      <el-table-column prop="id" label="编号" width="100" />
      <el-table-column prop="name" label="设备名称" min-width="160" />
      <el-table-column prop="location" label="位置" min-width="160" />
      <el-table-column prop="prizes" label="奖品数" width="90" align="center" />
      <el-table-column prop="coins" label="今日投币" width="100" align="center" />
      <el-table-column prop="revenue" label="今日收益" width="100" align="center">
        <template #default="{ row }">¥{{ row.revenue }}</template>
      </el-table-column>
      <el-table-column prop="status" label="状态" width="90" align="center">
        <template #default="{ row }">
          <el-tag :type="tagType(row.status)" size="small">{{ row.status }}</el-tag>
        </template>
      </el-table-column>
      <el-table-column label="操作" width="140" fixed="right">
        <template #default>
          <el-button text type="primary" size="small">详情</el-button>
          <el-button text type="warning" size="small">维护</el-button>
          <el-button text type="danger" size="small">下线</el-button>
        </template>
      </el-table-column>
    </el-table>

    <div style="margin-top:16px;display:flex;justify-content:flex-end">
      <el-pagination background layout="prev,pager,next" :total="machines.length" :page-size="10" />
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'

const search = ref('')
const filterStatus = ref('')

const machines = [
  { id: 'M-001', name: '商场A-01号机', location: '朝阳区国贸商场B1', prizes: 48, coins: 126, revenue: '630', status: '在线' },
  { id: 'M-002', name: '商场A-02号机', location: '朝阳区国贸商场B1', prizes: 32, coins: 98, revenue: '490', status: '在线' },
  { id: 'M-003', name: '商场B-01号机', location: '海淀区中关村广场', prizes: 12, coins: 45, revenue: '225', status: '维护' },
  { id: 'M-004', name: '商场B-03号机', location: '海淀区中关村广场', prizes: 60, coins: 210, revenue: '1050', status: '在线' },
  { id: 'M-005', name: '地铁站-01号机', location: '西城区西单地铁站', prizes: 5, coins: 0, revenue: '0', status: '离线' },
  { id: 'M-006', name: '商场C-01号机', location: '丰台区大红门商城', prizes: 40, coins: 88, revenue: '440', status: '在线' },
]

const filteredMachines = computed(() =>
  machines.filter(m => {
    const s = search.value.toLowerCase()
    const matchSearch = !s || m.name.includes(s) || m.id.toLowerCase().includes(s)
    const statusMap = { online: '在线', maintain: '维护', offline: '离线' }
    const matchStatus = !filterStatus.value || m.status === statusMap[filterStatus.value]
    return matchSearch && matchStatus
  })
)

function tagType(status) {
  return { '在线': 'success', '维护': 'warning', '离线': 'danger' }[status] ?? 'info'
}
</script>

<style scoped>
.toolbar { display: flex; gap: 12px; align-items: center; }
</style>
