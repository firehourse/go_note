
```mathematica
XGBoost 学习路线图 (@你)

├── 入门基础
│   ├── 什么是 Boosting
│   │   ├── 加法模型（Additive Model）
│   │   ├── 残差学习
│   │   └── GBDT → XGBoost 的改进
│   ├── XGBoost 定位
│   │   ├── 为什么比 GBDT 强
│   │   ├── 为什么训练快
│   │   └── 为什么更稳
│   └── 环境搭建
│       ├── Python + xgboost
│       ├── scikit-learn API
│       └── Jupyter / VSCode

├── 基础算法（核心原理）
│   ├── 二阶梯度（Gradient / Hessian）
│   ├── 树的增量训练流程（Boosting Round）
│   ├── 分裂公式（Gain）
│   ├── L1 / L2 正则化
│   ├── Shrinkage（Learning Rate）
│   └── 列抽样、行抽样

├── XGBoost 内部结构拆解（类似你研究 Go context）
│   ├── Booster（模型容器）
│   ├── DMatrix（输入数据结构）
│   ├── Tree Model（每棵树的节点结构）
│   ├── Objective（loss / grad / hess）
│   └── Training Flow（训练流程拆解）

├── 参数体系（最容易卡住的部分）
│   ├── 树结构参数
│   │   ├── max_depth
│   │   ├── min_child_weight
│   │   ├── gamma
│   │   └── reg_alpha / reg_lambda
│   ├── Boosting 类参数
│   │   ├── learning_rate
│   │   ├── n_estimators
│   │   └── booster (gbtree/dart)
│   ├── 随机性参数
│   │   ├── subsample
│   │   └── colsample_bytree
│   └── 训练控制
│       ├── eval_set
│       ├── early_stopping_rounds
│       └── verbose

├── 实验与调参（实践区）
│   ├── 单变量扫描
│   │   ├── 深度扫描（max_depth）
│   │   ├── 学习率扫描（learning_rate）
│   │   └── 采样率扫描（subsample）
│   ├── Overfitting vs Underfitting
│   ├── 训练/验证曲线观测
│   └── 模型版本记录（日志式）

├── 模型可解释性
│   ├── Feature Importance
│   │   ├── gain
│   │   ├── cover
│   │   └── weight
│   ├── Model Dump（解析一棵树）
│   └── SHAP（局部解释）

├── 实战资料集
│   ├── 分类任务（binary/multi）
│   ├── 回归任务
│   ├── Kaggle 数据
│   └── 自己的业务数据（你爬虫或 Log 数据）

└── 进阶拓展
    ├── LightGBM 对比
    ├── CatBoost 对比
    ├── XGBoost + 特征工程
    └── 模型部署（pickle / python API）
