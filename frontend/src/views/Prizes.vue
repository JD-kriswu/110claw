<template>
  <div>
    <div class="toolbar">
      <el-input v-model="search" placeholder="搜索奖品名称" clearable style="width:240px">
        <template #prefix><el-icon><Search /></el-icon></template>
      </el-input>
      <el-button type="primary" :icon="Plus">添加奖品</el-button>
    </div>

    <el-row :gutter="16" style="margin-top:16px">
      <el-col :span="6" v-for="p in filteredPrizes" :key="p.id" style="margin-bottom:16px">
        <div class="prize-card">
          <div class="prize-img" :style="{ background: p.color }">
            <el-icon size="40" color="white"><GiftFilled /></el-icon>
          </div>
          <div class="prize-info">
            <div class="prize-name">{{ p.name }}</div>
            <div class="prize-meta">
              <span class="prize-cost">¥{{ p.cost }}</span>
              <el-tag size="small" :type="p.stock > 10 ? 'success' : p.stock > 0 ? 'warning' : 'danger'">
                库存 {{ p.stock }}
              </el-tag>
            </div>
          </div>
          <div class="prize-actions">
            <el-button text type="primary" size="small">编辑</el-button>
            <el-button text type="danger" size="small">删除</el-button>
          </div>
        </div>
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { Plus } from '@element-plus/icons-vue'

const search = ref('')

const prizes = [
  { id: 1, name: '皮卡丘公仔', cost: '45', stock: 120, color: '#f59e0b' },
  { id: 2, name: '蜡笔小新玩偶', cost: '38', stock: 85, color: '#6366f1' },
  { id: 3, name: '熊本熊公仔', cost: '52', stock: 6, color: '#10b981' },
  { id: 4, name: '奥特曼变身器', cost: '29', stock: 0, color: '#ef4444' },
  { id: 5, name: '限定盲盒', cost: '99', stock: 30, color: '#8b5cf6' },
  { id: 6, name: '哆啦A梦抱枕', cost: '60', stock: 15, color: '#3b82f6' },
  { id: 7, name: '史努比钥匙扣', cost: '18', stock: 200, color: '#f97316' },
  { id: 8, name: '迷你波利卡车', cost: '35', stock: 3, color: '#14b8a6' },
]

const filteredPrizes = computed(() =>
  prizes.filter(p => !search.value || p.name.includes(search.value))
)
</script>

<style scoped>
.toolbar { display: flex; gap: 12px; align-items: center; }
.prize-card {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 1px 3px rgba(0,0,0,.06);
  transition: box-shadow .2s;
}
.prize-card:hover { box-shadow: 0 4px 12px rgba(0,0,0,.1); }
.prize-img {
  height: 120px;
  display: flex; align-items: center; justify-content: center;
}
.prize-info { padding: 12px 14px 0; }
.prize-name { font-size: 14px; font-weight: 600; color: #111827; margin-bottom: 8px; }
.prize-meta { display: flex; justify-content: space-between; align-items: center; }
.prize-cost { font-size: 16px; font-weight: 700; color: #6366f1; }
.prize-actions { display: flex; padding: 8px 14px 12px; gap: 4px; }
</style>
