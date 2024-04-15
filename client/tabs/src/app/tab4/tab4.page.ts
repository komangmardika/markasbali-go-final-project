import {Component, OnInit, ViewChild} from '@angular/core';
import {CommonService} from "./services/common.service";
import {BackupListInterface, SingleBackupListInterface} from "./models/backup-list.interface";
import {IonModal} from "@ionic/angular";
import { OverlayEventDetail } from '@ionic/core/components';
import {Subject, take, takeUntil} from "rxjs";
@Component({
  selector: 'app-tab4',
  templateUrl: 'tab4.page.html',
  styleUrls: ['tab4.page.scss']
})
export class Tab4Page implements OnInit{
  @ViewChild(IonModal, { static: true }) modal: IonModal | undefined;
  modalOpen = false;
  private destroy$: Subject<void> = new Subject<void>();
  backupList: BackupListInterface[] = [];
  histories: SingleBackupListInterface = {database_name: ''};
  dbName: string = '';
  constructor(
    private commonService: CommonService
  ) {}

  ngOnInit() {
    this.commonService.latestBackedUpAllDatabases().pipe(takeUntil(this.destroy$)).subscribe(r => {
      this.backupList = r.data as BackupListInterface[];
    }, err => {
      console.log(err);
    })
  }

  refresh() {
    this.commonService.latestBackedUpAllDatabases().pipe(takeUntil(this.destroy$)).subscribe(r => {
      this.backupList = r.data as BackupListInterface[];
    }, err => {
      console.log(err);
    })
  }

  cancel() {
    if(this.modal) {
      this.modal.dismiss(null, 'cancel');
      this.modalOpen = true;
    }
  }

  openDetail(bl: BackupListInterface) {
    this.commonService.allBackedUpOneDatabases(bl.database_name).pipe(takeUntil(this.destroy$)).subscribe(r => {
      this.dbName = bl.database_name;
      this.histories = r.data as SingleBackupListInterface;
      this.modalOpen = true;
    }, err => {
      console.log(err)
    })
  }

  onWillDismiss(event: Event) {
    const ev = event as CustomEvent<OverlayEventDetail<string>>;
    this.modalOpen = false;
  }

}
