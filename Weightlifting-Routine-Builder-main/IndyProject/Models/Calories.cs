using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace IndyProject.Models
{
    public class Calories
    {

        private double _bmr;
        private double _weight;
        private double _height;
        private int _age;
        private double _orm;
        private int _weightLifted;
        private int _reps;
        private double _lIntensityLoss;
        private double _mIntensityLoss;
        private double _hIntensityLoss;
        private double _lIntensityGain;
        private double _mIntensityGain;
        private double _hIntensityGain;

        public double bmr
        {
            get { return _bmr; }
            set
            {
                _bmr = value;
            }
        }
        public double weight // Weight in lbs
        {
            get { return _weight; }
            set
            {
                _weight = value;
            }
        }
        public double height // Height in centimeters
        {
            get { return _height; }
            set
            {
                _height = value;
            }
        }
        public int age // Age in years
        {
            get { return _age; }
            set
            {
                _age = value;
            }
        }
        public double orm
        {
            get { return _orm; }
            set
            {
                _orm = value;
            }
        }
        public int weightLifted
        {
            get { return _weightLifted; }
            set
            {
                _weightLifted = value;
            }
        }
        public int reps
        {
            get { return _reps; }
            set
            {
                _reps = value;
            }
        }
        public double lIntensityLoss
        {
            get { return _lIntensityLoss; }
            set
            {
                _lIntensityLoss = value;
            }
        }
        public double mIntensityLoss
        {
            get { return _mIntensityLoss; }
            set
            {
                _mIntensityLoss = value;
            }
        }
        public double hIntensityLoss
        {
            get { return _hIntensityLoss; }
            set
            {
                _hIntensityLoss = value;
            }
        }
        public double lIntensityGain
        {
            get { return _lIntensityGain; }
            set
            {
                _lIntensityGain = value;
            }
        }
        public double mIntensityGain
        {
            get { return _mIntensityGain; }
            set
            {
                _mIntensityGain = value;
            }
        }
        public double hIntensityGain
        {
            get { return _hIntensityGain; }
            set
            {
                _hIntensityGain = value;
            }
        }

        public void calcBMR()
        {
            if ((_height == 0) || (_weight == 0) || (_age == 0))
            {
                _bmr = 0;
            }
            else
            _bmr = (10 * _weight) + (6.25 * _height) - (5 * _age) + 5;
            _lIntensityLoss = _bmr - 250;
            _mIntensityLoss = _bmr - 500;
            _hIntensityLoss = _bmr - 750;
            _lIntensityGain = _bmr + 250;
            _mIntensityGain = _bmr + 500;
            _hIntensityGain = _bmr + 750;
        }
        public void calcORM()
        {
            _orm = _weightLifted / (1.0278 - 0.0278 * _reps );
        }
    }
}
