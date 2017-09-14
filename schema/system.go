package schema

var (
	system = `
	type System {
		info: SystemInfo
		version: SystemVersion
	}
	type SystemVersion {
		apiVersion: String!
		arch: String!
		buildTime: String!
		experimental: Boolean!
		gitCommit: String!
		goVersion: String!
		kernelVersion: String!
		minApiVersion: String!
		os: String!
		version: String!
	}
	type SystemInfo {		
		### TBI
		# containerdCommit: Commit
		# driverStatus: [[String]] # [][2]string
		# genericResources: [SwarmGenericResource]
		# initCommit: Commit
		# runcCommit: Commit
		# runtimes: [Runtime] # Returns a map[string]Runtime
		# systemStatus: [[String]] # [][2]string
	
		architecture: String
		bridgeNfIp6tables: Boolean
		bridgeNfIptables: Boolean
		cgroupDriver: String
		clusterAdvertise: String
		clusterStore: String
		containers: Int
		containersPaused: Int
		containersRunning: Int
		containersStopped: Int
		cpuCfsPeriod: Boolean
		cpuCfsQuota: Boolean
		cpuSet: Boolean
		cpuShares: Boolean
		debug: Boolean
		defaultRuntime: String
		dockerRootDir: String
		driver: String
		experimentalBuild: Boolean
		httpProxy: String
		httpsProxy: String
		id: ID!
		images: Int
		indexServerAddress: String
		initBinary: String
		ipv4Forwarding: Boolean
		# isolation: ContainerIsolation
		kernelMemory: Boolean
		kernelVersion: String
		labels: [String!]
		liveRestoreEnabled: Boolean
		loggingDriver: String
		memTotal: Int #This was a Long, but the struct uses an Int64
		memoryLimit: Boolean
		nCpu: Int
		nEventsListener: Int
		nFd: Int
		nGoRoutines: Int
		name: String
		noProxy: String
		oomKillDisable: Boolean
		operatingSystem: String
		osType: String
		# plugins: Plugins
		# registryConfig: RegistryConfig
		securityOptions: [String!]
		serverVersion: String
		swapLimit: Boolean
		# swarm: SwarmInfo
		systemTime: String
	}
	`
)
