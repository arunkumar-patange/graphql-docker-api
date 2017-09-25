package schema

var (
	swarm = `
	type Swarm {
	 	createdAt: String
	 	id: ID!
		joinTokens: JoinTokens
	# 	nodes(filter: NodeFilter): [SwarmNode]
	 	swarmSpec: SwarmSpec
	#	unlockKey: String #UnlockKey
	 	version: SwarmVersion
	}
	
	type SwarmSpec {
		# annotations: [Annotations]
	 	caConfig: CaConfig
	 	dispatcher: DispatcherConfig
	 	encryptionConfig: EncryptionConfig
	 	# TBI
	 	# labels: StringAnyMap
	 	name: String
	 	orchestration: OrchestrationConfig
	 	raft: RaftConfig
	 	taskDefaults: TaskDefaults
	}
	
	type OrchestrationConfig {
		taskHistoryRetentionLimit: Int
	}
	
	type RaftConfig {
		electionTick: Int
		heartbeatTick: Int
		keepOldSnapshots: Int
		logEntriesForSlowFollowers: Int
		snapshotInterval: Int
	}
	
	type DispatcherConfig {
	 	heartbeatPeriod: String
	}
	
	type CaConfig {
		externalCas: [ExternalCa]
		nodeCertExpiry: String
	}
	
	type ExternalCa {
		caCert: String
		# options: StringAnyMap
		protocol: String
		url: String
	}
	
	type EncryptionConfig {
		autoLockManagers: Boolean
	}
	
	type TaskDefaults {
	 	logDriver: Driver
	}
	
	type Driver {
		name: String
		# TBI
		# options: StringAnyMap
	}
	
	type JoinTokens {
		manager: String
		worker: String
	}
	
	# type UnlockKey {
	# 	unlockKey: String
	# }
	
	# type SwarmNode {
	# 	createdAt: Date
	# 	description: NodeDescription
	# 	details: NodeInfo
	# 	id: String
	# 	managerStatus: ManagerStatus
	# 	spec: NodeSpec
	# 	status: NodeStatus
	# 	updatedAt: Date
	# 	version: SwarmVersion
	# }
	
	# type NodeSpec {
	# 	availability: String
	# 	labels: StringAnyMap
	# 	name: String
	# 	role: String
	# }
	
	# type NodeDescription {
	# 	engine: EngineConfig
	# 	hostname: String
	# 	platform: Platform
	# 	resources: Resources
	# }
	
	# type Platform {
	# 	architecture: String
	# 	os: String
	# }
	
	# type Resources {
	# 	memoryBytes: Long
	# 	nanoCpus: Long
	# }
	
	# type EngineConfig {
	# 	engineVersion: String
	# 	labels: StringAnyMap
	# 	plugins: [EnginePlugin]
	# }
	
	# type EnginePlugin {
	# 	name: String
	# 	type: String
	# }
	
	# type NodeStatus {
	# 	addr: String
	# 	state: String
	# }
	
	# type ManagerStatus {
	# 	addr: String
	# 	leader: Boolean
	# 	reachability: String
	# }
	
	# input NodeFilter {
	# 	id: [String]
	# 	label: [String]
	# 	membership: [NodeMembership]
	# 	name: [String]
	# 	role: [NodeRole]
	# }
	
	# enum NodeMembership {
	# 	accepted
	# 	pending
	# }
	
	# enum NodeRole {
	# 	manager
	# 	worker
	# }
	
	# type NodeInfo {
	# 	createdAt: Date
	# 	description: NodeDescription
	# 	id: ID!
	# 	managerStatus: ManagerStatus
	# 	spec: NodeSpec
	# 	status: NodeStatus
	# 	updatedAt: Date
	# 	version: SwarmVersion
	# }
	
	type SwarmInfo {
		cluster: SwarmCluster
		controlAvailable: Boolean
		error: String
		localNodeState: String
		managers: Int
		nodeAddr: String
		nodeId: String
		nodes: Int
		remoteManagers: [RemoteManager]
	}
	
	type SwarmCluster {
		createdAt: String
	 	id: ID!
	 	rootRotationInProgress: Boolean
	 	swarmSpec: SwarmSpec
	 	tlsInfo: TLSInfo
	 	updatedAt: String
	 	version: SwarmVersion
	}
	
	type SwarmVersion {
		index: Int
	}

	type TLSInfo{
		certIssuerPublicKey: String
		certIssuerSubject: String
		trustRoot: String
	}
	
	type RemoteManager {
		addr: String
		nodeId: String
	}
	`
)
