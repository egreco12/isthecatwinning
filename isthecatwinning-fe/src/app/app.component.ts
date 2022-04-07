import { Component } from '@angular/core';
import { ScoreService } from './score.service'
import { Player } from './players';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'isthecatwinning';
  players: Player[] = [];
  firstPlace: Player = new Player();
  tigerWoods: Player = new Player();
  jordo: Player = new Player();
  differential: number = 0;

  constructor(private scoreService: ScoreService) {}

  tigerIsWinning(): boolean {
    return this.players[0].name === "Tiger Woods";	
  }


  ngOnInit(): void {
    // Current hardcoded order: 0 -> first place, 1 -> tiger woods, 2 -> jordo
    this.scoreService.getJSON().subscribe(
	res => {
	 console.log('HTTP response', res);
         res.forEach((player: Player)=> {
           this.players.push(player);
	 });
         this.firstPlace = this.players[0];
	 this.tigerWoods = this.players[1];
         this.jordo = this.players[2];
         this.differential = -1*(this.firstPlace.totalScore - this.tigerWoods.totalScore);
        },
	err => console.log('HTTP Error', err)
    );
  }
}
