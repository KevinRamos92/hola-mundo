#include <stdio.h>
#include <conio.h>
#include <iostream.h>
float A,B,H;
main()
{

      printf("Programa Que Calcula El Area De Un Triangulo\n");
      printf("Dame La Base\t");
      scanf("%f",&B);
      printf("Dame La Altura\t");
      scanf("%f",&H);
      A=B*H/2;
      printf("El Area Del Triangulo Es:  %.1f",A);
      getch();
      }