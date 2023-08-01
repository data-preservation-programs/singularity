# V1 或 V2

[Singularity V1](https://github.com/tech-greedy/singularity) 源于一个旨在解决有效将数据引入 Filecoin 存储提供商中的可扩展性的 DGM 研究项目。Tech Greedy 和 Kernelogic 接受了该研究项目的资助，并开发出了第一版 Singularity，逐渐成为 Filecoin 生态系统中最受欢迎的数据引入工具。

通过之前实施 V1 的经验，我们现在正在转向 V2，这是一个全新的实现，具有与 V1 相同的功能，并且带来更多具有颠覆性的功能。以下是 V1 和 V2 的简要比较：

<table>
  <thead>
    <tr>
      <th width="285.3333333333333">功能</th>
      <th>Singularity V1</th>
      <th>Singularity V2</th>
    </tr>
  </thead>
  <tbody>
    <tr>
      <td>语言</td>
      <td>Node.js + Golang</td>
      <td>Golang</td>
    </tr>
    <tr>
      <td>安装</td>
      <td>npm install</td>
      <td>go install</td>
    </tr>
    <tr>
      <td>本地文件系统数据源</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>远程数据源</td>
      <td>公共 S3 存储桶</td>
      <td>40+ 集成</td>
    </tr>
    <tr>
      <td>上传 / 推送新文件</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>内联准备</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>加密</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>维护数据集层次结构</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>CAR 下载</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>文件下载</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>交易制作调度器</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>交易制作自助服务</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>零外部依赖</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>钱包管理</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>远程签名</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
    <tr>
      <td>Web API</td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="274c">❌</span></td>
      <td><span data-gb-custom-inline data-tag="emoji" data-code="2705">✅</span></td>
    </tr>
  </tbody>
</table>