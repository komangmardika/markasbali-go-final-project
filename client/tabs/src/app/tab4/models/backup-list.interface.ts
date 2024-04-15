export interface BackupListInterface {
  database_name : string;
  latest_backup ?: LatestBackupInterface;
}

export interface SingleBackupListInterface {
  database_name ?: string;
  histories ?: LatestBackupInterface[];
}

export interface LatestBackupInterface {
  id ?: number;
  file_name ?: string;
  timestamp ?: string;
}
