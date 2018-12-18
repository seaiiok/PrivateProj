CREATE TABLE [dbo].[iFixsvr_JFofflinefiles](
	[Id] [decimal](18, 0) IDENTITY(1,1) NOT NULL,
	[J_DeviceName] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_DeviceIP] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_FileName] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_FilePath] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_FileCreateDT] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_GoodSum] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_BadSum] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_StdNo] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_PKNo] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_Coder] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL,
	[J_Dout] [nvarchar](50) COLLATE Chinese_PRC_CI_AS NULL
) ON [PRIMARY]
